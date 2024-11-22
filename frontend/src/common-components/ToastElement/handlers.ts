import {showCloseToastAnimation} from "@/common-components/ToastElement/helpers.ts";

export const handleToastCloseButtonClick = () => {
    const timeout = showCloseToastAnimation();
    return () => clearTimeout(timeout);
}