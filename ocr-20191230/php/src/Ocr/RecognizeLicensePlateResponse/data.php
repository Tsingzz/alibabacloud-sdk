<?php

// This file is auto-generated, don't edit it. Thanks.

namespace AlibabaCloud\SDK\Ocr\V20191230\Ocr\RecognizeLicensePlateResponse;

use AlibabaCloud\Tea\Model;

class data extends Model
{
    /**
     * @description plates
     *
     * @var array
     */
    public $plates;
    protected $_name = [
        'plates' => 'Plates',
    ];

    public function validate()
    {
        Model::validateRequired('plates', $this->plates, true);
    }

    public function toMap()
    {
        $res           = [];
        $res['Plates'] = [];
        if (null !== $this->plates && \is_array($this->plates)) {
            $n = 0;
            foreach ($this->plates as $item) {
                $res['Plates'][$n++] = null !== $item ? $item->toMap() : $item;
            }
        }

        return $res;
    }

    /**
     * @param array $map
     *
     * @return data
     */
    public static function fromMap($map = [])
    {
        $model = new self();
        if (isset($map['Plates'])) {
            if (!empty($map['Plates'])) {
                $model->plates = [];
                $n             = 0;
                foreach ($map['Plates'] as $item) {
                    $model->plates[$n++] = null !== $item ? RecognizeLicensePlateResponse\data\plates::fromMap($item) : $item;
                }
            }
        }

        return $model;
    }
}
