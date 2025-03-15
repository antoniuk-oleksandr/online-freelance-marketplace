import SelectMenuContent
    from "@/common-components/Select/components/SelectMenuContent/SelectMenuContent.svelte";
import type {SelectMenuProps} from "@/types/SelectMenuProps";
import {mount} from "svelte";

export const resolveSelectAnimation = (
    timeout: number | undefined,
    setTimeoutValue: (value: number | undefined) => void,
    isOpen: boolean,
    setShowExitAnimation: (value: boolean) => void,
) => {
    clearTimeout(timeout);

    if (isOpen) setShowExitAnimation(true);
    else {
        const value = setTimeout(() => {
            setShowExitAnimation(false);
        }, 300);

        setTimeoutValue(value);
    }
}

export const getSelectModal = (
    SelectMenuProps: SelectMenuProps
) => {
    return {
        render: (targetElement: HTMLElement) => {
            mount(SelectMenuContent, {
                target: targetElement,
                props: SelectMenuProps
            })
        }
    };
};
