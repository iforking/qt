package nfc

//#include "nfc.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QNearFieldShareManager struct {
	core.QObject
}

type QNearFieldShareManager_ITF interface {
	core.QObject_ITF
	QNearFieldShareManager_PTR() *QNearFieldShareManager
}

func PointerFromQNearFieldShareManager(ptr QNearFieldShareManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNearFieldShareManager_PTR().Pointer()
	}
	return nil
}

func NewQNearFieldShareManagerFromPointer(ptr unsafe.Pointer) *QNearFieldShareManager {
	var n = new(QNearFieldShareManager)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QNearFieldShareManager_") {
		n.SetObjectName("QNearFieldShareManager_" + qt.Identifier())
	}
	return n
}

func (ptr *QNearFieldShareManager) QNearFieldShareManager_PTR() *QNearFieldShareManager {
	return ptr
}

//QNearFieldShareManager::ShareError
type QNearFieldShareManager__ShareError int64

const (
	QNearFieldShareManager__NoError                     = QNearFieldShareManager__ShareError(0)
	QNearFieldShareManager__UnknownError                = QNearFieldShareManager__ShareError(1)
	QNearFieldShareManager__InvalidShareContentError    = QNearFieldShareManager__ShareError(2)
	QNearFieldShareManager__ShareCanceledError          = QNearFieldShareManager__ShareError(3)
	QNearFieldShareManager__ShareInterruptedError       = QNearFieldShareManager__ShareError(4)
	QNearFieldShareManager__ShareRejectedError          = QNearFieldShareManager__ShareError(5)
	QNearFieldShareManager__UnsupportedShareModeError   = QNearFieldShareManager__ShareError(6)
	QNearFieldShareManager__ShareAlreadyInProgressError = QNearFieldShareManager__ShareError(7)
	QNearFieldShareManager__SharePermissionDeniedError  = QNearFieldShareManager__ShareError(8)
)

//QNearFieldShareManager::ShareMode
type QNearFieldShareManager__ShareMode int64

const (
	QNearFieldShareManager__NoShare   = QNearFieldShareManager__ShareMode(0x00)
	QNearFieldShareManager__NdefShare = QNearFieldShareManager__ShareMode(0x01)
	QNearFieldShareManager__FileShare = QNearFieldShareManager__ShareMode(0x02)
)

func NewQNearFieldShareManager(parent core.QObject_ITF) *QNearFieldShareManager {
	defer qt.Recovering("QNearFieldShareManager::QNearFieldShareManager")

	return NewQNearFieldShareManagerFromPointer(C.QNearFieldShareManager_NewQNearFieldShareManager(core.PointerFromQObject(parent)))
}

func (ptr *QNearFieldShareManager) ConnectError(f func(error QNearFieldShareManager__ShareError)) {
	defer qt.Recovering("connect QNearFieldShareManager::error")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_ConnectError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error", f)
	}
}

func (ptr *QNearFieldShareManager) DisconnectError() {
	defer qt.Recovering("disconnect QNearFieldShareManager::error")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_DisconnectError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error")
	}
}

//export callbackQNearFieldShareManagerError
func callbackQNearFieldShareManagerError(ptrName *C.char, error C.int) {
	defer qt.Recovering("callback QNearFieldShareManager::error")

	var signal = qt.GetSignal(C.GoString(ptrName), "error")
	if signal != nil {
		signal.(func(QNearFieldShareManager__ShareError))(QNearFieldShareManager__ShareError(error))
	}

}

func (ptr *QNearFieldShareManager) SetShareModes(mode QNearFieldShareManager__ShareMode) {
	defer qt.Recovering("QNearFieldShareManager::setShareModes")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_SetShareModes(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QNearFieldShareManager) ShareError() QNearFieldShareManager__ShareError {
	defer qt.Recovering("QNearFieldShareManager::shareError")

	if ptr.Pointer() != nil {
		return QNearFieldShareManager__ShareError(C.QNearFieldShareManager_ShareError(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNearFieldShareManager) ShareModes() QNearFieldShareManager__ShareMode {
	defer qt.Recovering("QNearFieldShareManager::shareModes")

	if ptr.Pointer() != nil {
		return QNearFieldShareManager__ShareMode(C.QNearFieldShareManager_ShareModes(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNearFieldShareManager) ConnectShareModesChanged(f func(modes QNearFieldShareManager__ShareMode)) {
	defer qt.Recovering("connect QNearFieldShareManager::shareModesChanged")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_ConnectShareModesChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "shareModesChanged", f)
	}
}

func (ptr *QNearFieldShareManager) DisconnectShareModesChanged() {
	defer qt.Recovering("disconnect QNearFieldShareManager::shareModesChanged")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_DisconnectShareModesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "shareModesChanged")
	}
}

//export callbackQNearFieldShareManagerShareModesChanged
func callbackQNearFieldShareManagerShareModesChanged(ptrName *C.char, modes C.int) {
	defer qt.Recovering("callback QNearFieldShareManager::shareModesChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "shareModesChanged")
	if signal != nil {
		signal.(func(QNearFieldShareManager__ShareMode))(QNearFieldShareManager__ShareMode(modes))
	}

}

func QNearFieldShareManager_SupportedShareModes() QNearFieldShareManager__ShareMode {
	defer qt.Recovering("QNearFieldShareManager::supportedShareModes")

	return QNearFieldShareManager__ShareMode(C.QNearFieldShareManager_QNearFieldShareManager_SupportedShareModes())
}

func (ptr *QNearFieldShareManager) ConnectTargetDetected(f func(shareTarget *QNearFieldShareTarget)) {
	defer qt.Recovering("connect QNearFieldShareManager::targetDetected")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_ConnectTargetDetected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "targetDetected", f)
	}
}

func (ptr *QNearFieldShareManager) DisconnectTargetDetected() {
	defer qt.Recovering("disconnect QNearFieldShareManager::targetDetected")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_DisconnectTargetDetected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "targetDetected")
	}
}

//export callbackQNearFieldShareManagerTargetDetected
func callbackQNearFieldShareManagerTargetDetected(ptrName *C.char, shareTarget unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldShareManager::targetDetected")

	var signal = qt.GetSignal(C.GoString(ptrName), "targetDetected")
	if signal != nil {
		signal.(func(*QNearFieldShareTarget))(NewQNearFieldShareTargetFromPointer(shareTarget))
	}

}

func (ptr *QNearFieldShareManager) DestroyQNearFieldShareManager() {
	defer qt.Recovering("QNearFieldShareManager::~QNearFieldShareManager")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_DestroyQNearFieldShareManager(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}