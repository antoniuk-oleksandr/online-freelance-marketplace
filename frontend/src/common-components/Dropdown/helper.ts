import DropdownMenuContent
    from "@/common-components/Dropdown/components/DropdownMenuContent/DropdownMenuContent.svelte";
import type {DropdownMenuProps} from "@/types/DropdownMenuProps.ts";
import {mount} from "svelte";

export const resolveDropdownAnimation = (
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

export const getDropdownModal = (
    dropdownMenuProps: DropdownMenuProps
) => {
    return {
        render: (targetElement: HTMLElement) => {
            mount(DropdownMenuContent, {
                target: targetElement,
                props: dropdownMenuProps
            })
        }
    };
};