package _4_obj_pool

import (
	"errors"
	"testing"
	"time"
)

type ReusableObj struct {
}

type ObjPool struct {
	bufChan chan *ReusableObj
}

func NewObjPool(numOfObj int) *ObjPool {
	op := ObjPool{}
	op.bufChan = make(chan *ReusableObj, numOfObj)
	for i := 0; i < numOfObj; i++ {
		op.bufChan <- &ReusableObj{}
	}
	return &op
}

func (o *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error) {
	select {
	case ret := <-o.bufChan:
		return ret, nil
	case <-time.After(timeout):
		return nil, errors.New("timeout")
	}
}

func (o *ObjPool) ReleaseObj(obj *ReusableObj) error {
	select {
	case o.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}

func TestObjPool(t *testing.T) {
	pool := NewObjPool(10)

	for i := 0; i < 11; i++ {
		if obj, err := pool.GetObj(time.Second * 1); err != nil {
			t.Error(err)
		} else {
			t.Logf("%T\n", obj)
			if err := pool.ReleaseObj(obj); err != nil {
				t.Error(err)
			}
		}
	}
	t.Log("Done")
}
