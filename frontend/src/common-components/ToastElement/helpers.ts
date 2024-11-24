import {toastElementStore} from "@/common-components/ToastElement/store/toast-element-store.ts";

export const showOpenToastAnimation = () => setTimeout(() => {
    const timeout = showCloseToastAnimation();
    return () => clearTimeout(timeout);
}, 3000);

export const showCloseToastAnimation = () => {
    toastElementStore.update((prev) => ({
        ...prev,
        exitAnimation: true
    }));

    return  setTimeout(() => {
        toastElementStore.update((prev) => ({
            ...prev,
            show: false,
            exitAnimation: false
        }));
    }, 500);
}