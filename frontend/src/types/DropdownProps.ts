import type {LayoutProps} from "@/types/LayoutProps.ts";
import type {DropdownItem} from "@/types/DropdownItem.ts";

export type DropdownProps = LayoutProps & {
    items: DropdownItem[],
    selectedItem: string,
    title: string,
    additionalItems?: DropdownItem[],
    setSelectedItem?: (item: string) => void,
    selectedAdditionalItem?: string,
    setSelectedAdditionalItem?: (item: string) => void,
};