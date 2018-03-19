<?php

namespace Swaggest\JsonCli;

use Swaggest\JsonDiff\Exception;
use Swaggest\JsonDiff\JsonPatch;
use Yaoi\Command;
use Yaoi\Command\Definition;

class Apply extends Base
{
    public $patchPath;
    public $basePath;
    public $isURIFragmentId = false;
    public $tolerateErrors;

    /**
     * @param Definition $definition
     * @param \stdClass|static $options
     */
    static function setUpDefinition(Definition $definition, $options)
    {
        $options->patchPath = Command\Option::create()->setType()->setIsUnnamed()
            ->setDescription('Path to JSON patch file');
        $options->basePath = Command\Option::create()->setType()->setIsUnnamed()
            ->setDescription('Path to JSON base file');
        parent::setUpDefinition($definition, $options);
        $definition->description = 'Apply patch to base json document, output to STDOUT';
        $options->tolerateErrors = Command\Option::create()
            ->setDescription('Continue on error');
        unset($options->originalPath);
        unset($options->newPath);

    }

    public function performAction()
    {
        if (!file_exists($this->patchPath)) {
            $this->response->error('File not found: ' . $this->patchPath);
            return;
        }
        if (!file_exists($this->basePath)) {
            $this->response->error('File not found: ' . $this->basePath);
            return;
        }

        $patchJson = file_get_contents($this->patchPath);
        if (!$patchJson) {
            $this->response->error('Unable to read ' . $this->patchPath);
            return;
        }

        $baseJson = file_get_contents($this->basePath);
        if (!$baseJson) {
            $this->response->error('Unable to read ' . $this->basePath);
            return;
        }

        try {
            $patch = JsonPatch::import(json_decode($patchJson));
            $base = json_decode($baseJson);
            $errors = $patch->apply($base, !$this->tolerateErrors);
            foreach ($errors as $error) {
                $this->response->error($error->getMessage());
            }
            $this->out = $base;
        } catch (Exception $e) {
            $this->response->error($e->getMessage());
            die(1);
        }

        $this->postPerform();
    }

}