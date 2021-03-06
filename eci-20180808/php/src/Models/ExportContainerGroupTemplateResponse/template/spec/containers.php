<?php

// This file is auto-generated, don't edit it. Thanks.

namespace AlibabaCloud\SDK\Eci\V20180808\Models\ExportContainerGroupTemplateResponse\template\spec;

use AlibabaCloud\SDK\Eci\V20180808\Models\ExportContainerGroupTemplateResponse\template\spec\containers\env;
use AlibabaCloud\SDK\Eci\V20180808\Models\ExportContainerGroupTemplateResponse\template\spec\containers\livenessProbe;
use AlibabaCloud\SDK\Eci\V20180808\Models\ExportContainerGroupTemplateResponse\template\spec\containers\ports;
use AlibabaCloud\SDK\Eci\V20180808\Models\ExportContainerGroupTemplateResponse\template\spec\containers\readinessProbe;
use AlibabaCloud\SDK\Eci\V20180808\Models\ExportContainerGroupTemplateResponse\template\spec\containers\resources;
use AlibabaCloud\SDK\Eci\V20180808\Models\ExportContainerGroupTemplateResponse\template\spec\containers\securityContext;
use AlibabaCloud\SDK\Eci\V20180808\Models\ExportContainerGroupTemplateResponse\template\spec\containers\volumeMounts;
use AlibabaCloud\Tea\Model;

class containers extends Model
{
    /**
     * @description name
     *
     * @var string
     */
    public $name;

    /**
     * @description image
     *
     * @var string
     */
    public $image;

    /**
     * @description imagePullPolicy
     *
     * @var string
     */
    public $imagePullPolicy;

    /**
     * @description stdin
     *
     * @var bool
     */
    public $stdin;

    /**
     * @description stdinOnce
     *
     * @var bool
     */
    public $stdinOnce;

    /**
     * @description tty
     *
     * @var bool
     */
    public $tty;

    /**
     * @description workingDir
     *
     * @var string
     */
    public $workingDir;

    /**
     * @description env
     *
     * @var array
     */
    public $env;

    /**
     * @description ports
     *
     * @var array
     */
    public $ports;

    /**
     * @description volumeMounts
     *
     * @var array
     */
    public $volumeMounts;

    /**
     * @description securityContext
     *
     * @var securityContext
     */
    public $securityContext;

    /**
     * @description resources
     *
     * @var resources
     */
    public $resources;

    /**
     * @description readinessProbe
     *
     * @var readinessProbe
     */
    public $readinessProbe;

    /**
     * @description livenessProbe
     *
     * @var livenessProbe
     */
    public $livenessProbe;

    /**
     * @description command
     *
     * @var array
     */
    public $command;

    /**
     * @description args
     *
     * @var array
     */
    public $args;
    protected $_name = [
        'name'            => 'Name',
        'image'           => 'Image',
        'imagePullPolicy' => 'ImagePullPolicy',
        'stdin'           => 'Stdin',
        'stdinOnce'       => 'StdinOnce',
        'tty'             => 'Tty',
        'workingDir'      => 'WorkingDir',
        'env'             => 'Env',
        'ports'           => 'Ports',
        'volumeMounts'    => 'VolumeMounts',
        'securityContext' => 'SecurityContext',
        'resources'       => 'Resources',
        'readinessProbe'  => 'ReadinessProbe',
        'livenessProbe'   => 'LivenessProbe',
        'command'         => 'Command',
        'args'            => 'Args',
    ];

    public function validate()
    {
        Model::validateRequired('name', $this->name, true);
        Model::validateRequired('image', $this->image, true);
        Model::validateRequired('imagePullPolicy', $this->imagePullPolicy, true);
        Model::validateRequired('stdin', $this->stdin, true);
        Model::validateRequired('stdinOnce', $this->stdinOnce, true);
        Model::validateRequired('tty', $this->tty, true);
        Model::validateRequired('workingDir', $this->workingDir, true);
        Model::validateRequired('env', $this->env, true);
        Model::validateRequired('ports', $this->ports, true);
        Model::validateRequired('volumeMounts', $this->volumeMounts, true);
        Model::validateRequired('securityContext', $this->securityContext, true);
        Model::validateRequired('resources', $this->resources, true);
        Model::validateRequired('readinessProbe', $this->readinessProbe, true);
        Model::validateRequired('livenessProbe', $this->livenessProbe, true);
        Model::validateRequired('command', $this->command, true);
        Model::validateRequired('args', $this->args, true);
    }

    public function toMap()
    {
        $res = [];
        if (null !== $this->name) {
            $res['Name'] = $this->name;
        }
        if (null !== $this->image) {
            $res['Image'] = $this->image;
        }
        if (null !== $this->imagePullPolicy) {
            $res['ImagePullPolicy'] = $this->imagePullPolicy;
        }
        if (null !== $this->stdin) {
            $res['Stdin'] = $this->stdin;
        }
        if (null !== $this->stdinOnce) {
            $res['StdinOnce'] = $this->stdinOnce;
        }
        if (null !== $this->tty) {
            $res['Tty'] = $this->tty;
        }
        if (null !== $this->workingDir) {
            $res['WorkingDir'] = $this->workingDir;
        }
        if (null !== $this->env) {
            $res['Env'] = [];
            if (null !== $this->env && \is_array($this->env)) {
                $n = 0;
                foreach ($this->env as $item) {
                    $res['Env'][$n++] = null !== $item ? $item->toMap() : $item;
                }
            }
        }
        if (null !== $this->ports) {
            $res['Ports'] = [];
            if (null !== $this->ports && \is_array($this->ports)) {
                $n = 0;
                foreach ($this->ports as $item) {
                    $res['Ports'][$n++] = null !== $item ? $item->toMap() : $item;
                }
            }
        }
        if (null !== $this->volumeMounts) {
            $res['VolumeMounts'] = [];
            if (null !== $this->volumeMounts && \is_array($this->volumeMounts)) {
                $n = 0;
                foreach ($this->volumeMounts as $item) {
                    $res['VolumeMounts'][$n++] = null !== $item ? $item->toMap() : $item;
                }
            }
        }
        if (null !== $this->securityContext) {
            $res['SecurityContext'] = null !== $this->securityContext ? $this->securityContext->toMap() : null;
        }
        if (null !== $this->resources) {
            $res['Resources'] = null !== $this->resources ? $this->resources->toMap() : null;
        }
        if (null !== $this->readinessProbe) {
            $res['ReadinessProbe'] = null !== $this->readinessProbe ? $this->readinessProbe->toMap() : null;
        }
        if (null !== $this->livenessProbe) {
            $res['LivenessProbe'] = null !== $this->livenessProbe ? $this->livenessProbe->toMap() : null;
        }
        if (null !== $this->command) {
            $res['Command'] = [];
            if (null !== $this->command) {
                $res['Command'] = $this->command;
            }
        }
        if (null !== $this->args) {
            $res['Args'] = [];
            if (null !== $this->args) {
                $res['Args'] = $this->args;
            }
        }

        return $res;
    }

    /**
     * @param array $map
     *
     * @return containers
     */
    public static function fromMap($map = [])
    {
        $model = new self();
        if (isset($map['Name'])) {
            $model->name = $map['Name'];
        }
        if (isset($map['Image'])) {
            $model->image = $map['Image'];
        }
        if (isset($map['ImagePullPolicy'])) {
            $model->imagePullPolicy = $map['ImagePullPolicy'];
        }
        if (isset($map['Stdin'])) {
            $model->stdin = $map['Stdin'];
        }
        if (isset($map['StdinOnce'])) {
            $model->stdinOnce = $map['StdinOnce'];
        }
        if (isset($map['Tty'])) {
            $model->tty = $map['Tty'];
        }
        if (isset($map['WorkingDir'])) {
            $model->workingDir = $map['WorkingDir'];
        }
        if (isset($map['Env'])) {
            if (!empty($map['Env'])) {
                $model->env = [];
                $n          = 0;
                foreach ($map['Env'] as $item) {
                    $model->env[$n++] = null !== $item ? env::fromMap($item) : $item;
                }
            }
        }
        if (isset($map['Ports'])) {
            if (!empty($map['Ports'])) {
                $model->ports = [];
                $n            = 0;
                foreach ($map['Ports'] as $item) {
                    $model->ports[$n++] = null !== $item ? ports::fromMap($item) : $item;
                }
            }
        }
        if (isset($map['VolumeMounts'])) {
            if (!empty($map['VolumeMounts'])) {
                $model->volumeMounts = [];
                $n                   = 0;
                foreach ($map['VolumeMounts'] as $item) {
                    $model->volumeMounts[$n++] = null !== $item ? volumeMounts::fromMap($item) : $item;
                }
            }
        }
        if (isset($map['SecurityContext'])) {
            $model->securityContext = securityContext::fromMap($map['SecurityContext']);
        }
        if (isset($map['Resources'])) {
            $model->resources = resources::fromMap($map['Resources']);
        }
        if (isset($map['ReadinessProbe'])) {
            $model->readinessProbe = readinessProbe::fromMap($map['ReadinessProbe']);
        }
        if (isset($map['LivenessProbe'])) {
            $model->livenessProbe = livenessProbe::fromMap($map['LivenessProbe']);
        }
        if (isset($map['Command'])) {
            if (!empty($map['Command'])) {
                $model->command = [];
                $model->command = $map['Command'];
            }
        }
        if (isset($map['Args'])) {
            if (!empty($map['Args'])) {
                $model->args = [];
                $model->args = $map['Args'];
            }
        }

        return $model;
    }
}
