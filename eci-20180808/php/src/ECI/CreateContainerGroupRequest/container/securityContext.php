<?php

// This file is auto-generated, don't edit it. Thanks.

namespace AlibabaCloud\SDK\ECI\V20180808\ECI\CreateContainerGroupRequest\container;

use AlibabaCloud\Tea\Model;

class securityContext extends Model
{
    public $capability;

    public $readOnlyRootFilesystem;

    public $runAsUser;
    protected $_name = [
        'capability'             => 'Capability',
        'readOnlyRootFilesystem' => 'ReadOnlyRootFilesystem',
        'runAsUser'              => 'RunAsUser',
    ];
}
