<?php

// This file is auto-generated, don't edit it. Thanks.

namespace AlibabaCloud\SDK\Facebody\V20191230\Models\RecognizePublicFaceResponse\data\elements;

use AlibabaCloud\SDK\Facebody\V20191230\Models\RecognizePublicFaceResponse\data\elements\results\subResults;
use AlibabaCloud\Tea\Model;

class results extends Model
{
    /**
     * @var string
     */
    public $label;

    /**
     * @var string
     */
    public $suggestion;

    /**
     * @var float
     */
    public $rate;

    /**
     * @var array
     */
    public $subResults;
    protected $_name = [
        'label'      => 'Label',
        'suggestion' => 'Suggestion',
        'rate'       => 'Rate',
        'subResults' => 'SubResults',
    ];

    public function validate()
    {
        Model::validateRequired('label', $this->label, true);
        Model::validateRequired('suggestion', $this->suggestion, true);
        Model::validateRequired('rate', $this->rate, true);
        Model::validateRequired('subResults', $this->subResults, true);
    }

    public function toMap()
    {
        $res = [];
        if (null !== $this->label) {
            $res['Label'] = $this->label;
        }
        if (null !== $this->suggestion) {
            $res['Suggestion'] = $this->suggestion;
        }
        if (null !== $this->rate) {
            $res['Rate'] = $this->rate;
        }
        if (null !== $this->subResults) {
            $res['SubResults'] = [];
            if (null !== $this->subResults && \is_array($this->subResults)) {
                $n = 0;
                foreach ($this->subResults as $item) {
                    $res['SubResults'][$n++] = null !== $item ? $item->toMap() : $item;
                }
            }
        }

        return $res;
    }

    /**
     * @param array $map
     *
     * @return results
     */
    public static function fromMap($map = [])
    {
        $model = new self();
        if (isset($map['Label'])) {
            $model->label = $map['Label'];
        }
        if (isset($map['Suggestion'])) {
            $model->suggestion = $map['Suggestion'];
        }
        if (isset($map['Rate'])) {
            $model->rate = $map['Rate'];
        }
        if (isset($map['SubResults'])) {
            if (!empty($map['SubResults'])) {
                $model->subResults = [];
                $n                 = 0;
                foreach ($map['SubResults'] as $item) {
                    $model->subResults[$n++] = null !== $item ? subResults::fromMap($item) : $item;
                }
            }
        }

        return $model;
    }
}
