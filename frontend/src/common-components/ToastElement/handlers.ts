import { toastElementStore } from "./store/toast-element-store"

export const handleToastCloseButtonClick = () => {
    toastElementStore.update((prev) => ({ ...prev, show: false }))
}
