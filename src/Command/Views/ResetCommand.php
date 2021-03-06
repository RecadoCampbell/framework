<?php
/**
 * Spiral Framework.
 *
 * @license   MIT
 * @author    Anton Titov (Wolfy-J)
 */

namespace Spiral\Command\Views;

use Spiral\Console\Command;
use Spiral\Files\FilesInterface;
use Spiral\Views\Config\ViewsConfig;

/**
 * Remove every file located in view cache directory.
 */
class ResetCommand extends Command
{
    const NAME        = 'views:reset';
    const DESCRIPTION = 'Clear view cache';

    /**
     * @param ViewsConfig    $config
     * @param FilesInterface $files
     */
    public function perform(ViewsConfig $config, FilesInterface $files)
    {
        if (!$files->exists($config->cacheDirectory())) {
            $this->writeln("Cache directory is missing, no cache to be cleaned.");

            return;
        }

        if ($this->isVerbose()) {
            $this->writeln("<info>Cleaning view cache:</info>");
        }

        foreach ($files->getFiles($config->cacheDirectory()) as $filename) {
            try {
                $files->delete($filename);
            } catch (\Throwable $e) {
                $this->sprintf("<fg=red>[errored]</fg=red> `%s`: <fg=red>%s</fg=red>\n",
                    $files->relativePath($filename, $config->cacheDirectory()),
                    $e->getMessage()
                );

                continue;
            }

            if ($this->isVerbose()) {
                $this->sprintf(
                    "<fg=green>[deleted]</fg=green> `%s`\n",
                    $files->relativePath($filename, $config->cacheDirectory())
                );
            }
        }

        $this->writeln("<info>View cache has been cleared.</info>");
    }
}