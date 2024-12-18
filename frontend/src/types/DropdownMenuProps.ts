import {DropdownProps} from "@/types/DropdownProps.ts";

export type DropdownMenuProps = DropdownProps & {
    showExitAnimation: boolean,
    isOpen?: boolean,
    setIsOpen?: (value: boolean) => void,
};
