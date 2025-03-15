import type { SelectProps } from "./SelectProps";

export type SelectMenuProps = SelectProps & {
    showExitAnimation: boolean,
    isOpen?: boolean,
    setIsOpen?: (value: boolean) => void,
};
