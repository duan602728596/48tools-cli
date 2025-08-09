import { join } from 'node:path';
import { cwd } from 'node:process';
import { existsSync } from 'node:fs';
import { mkdir, copyFile, writeFile } from 'node:fs/promises';
import { promisify } from 'node:util';
import { execa } from 'execa';
import { zip } from 'cross-zip';
import packageJson from '../package.json' with { type: 'json' };

const zipAsync: (inPath: string, outPath: string) => Promise<void> = promisify(zip);

const softwareName: string = '48tools';
const osList: Array<string> = ['windows', 'linux', 'darwin'];
const archList: Array<string> = ['amd64', 'arm64'];
const buildDir: string = join(cwd(), '.build');

/**
 * 编译文件
 * @param { string } targetOs - 目标系统
 * @param { string } targetArch - 目标架构
 */
async function build(targetOs: string, targetArch: string): Promise<void> {
  // 创建目录
  const targetDir: string = join(buildDir, `${ targetOs }-${ targetArch }`);

  if (!existsSync(targetDir)) await mkdir(targetDir, { recursive: true });

  // 编译文件
  const isWindows: boolean = targetOs === 'windows';
  const outputFile: string = join(targetDir, isWindows ? `${ softwareName }.exe` : softwareName);

  console.log(`🚧 Building for ${ targetOs }/${ targetArch }...`);

  try {
    // 拷贝其他文件
    await Promise.all([
      execa('go', ['build', '-o', outputFile, 'src/main.go'], {
        env: {
          GOOS: targetOs,
          GOARCH: targetArch
        }
      }),
      copyFile(join(cwd(), 'README.md'), join(targetDir, 'README.md')),
      copyFile(join(cwd(), 'LICENSE'), join(targetDir, 'LICENSE')),
      copyFile(join(cwd(), 'config.example.yaml'), join(targetDir, 'config.yaml')),
      writeFile(join(targetDir, `v${ packageJson.version }`), '', { encoding: 'utf8' })
    ]);
    await zipAsync(targetDir, `${ packageJson.name }-${ packageJson.version }-${ targetOs }-${ targetArch }.zip`);

    console.log(`✅ Success: ${ outputFile }`);
  } catch (err) {
    console.error(`❌ Failed for ${ targetOs }/${ targetArch }:\n${ err }`);
  }
}

const promises: Array<PromiseLike<void>> = [];

for (const os of osList) {
  for (const arch of archList) {
    promises.push(build(os, arch));
  }
}

await Promise.all(promises);