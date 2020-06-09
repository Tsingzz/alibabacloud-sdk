<?php

// This file is auto-generated, don't edit it. Thanks.
namespace AlibabaCloud\SDK\Alidns\V20150109\Models;

use AlibabaCloud\Tea\Model;

class DescribeGtmAccessStrategyRequest extends Model {
    protected $_name = [
        'lang' => 'Lang',
        'strategyId' => 'StrategyId',
    ];
    public function validate() {
        Model::validateRequired('strategyId', $this->strategyId, true);
    }
    public function toMap() {
        $res = [];
        $res['Lang'] = $this->lang;
        $res['StrategyId'] = $this->strategyId;
        return $res;
    }
    /**
     * @param array $map
     * @return DescribeGtmAccessStrategyRequest
     */
    public static function fromMap($map = []) {
        $model = new self();
        if(isset($map['Lang'])){
            $model->lang = $map['Lang'];
        }
        if(isset($map['StrategyId'])){
            $model->strategyId = $map['StrategyId'];
        }
        return $model;
    }
    /**
     * @description lang
     * @var string
     */
    public $lang;

    /**
     * @description strategyId
     * @var string
     */
    public $strategyId;

}