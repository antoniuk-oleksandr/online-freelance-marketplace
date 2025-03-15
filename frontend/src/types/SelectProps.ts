import type { LayoutProps } from "@/types/LayoutProps.ts";
import type { SelectItem } from "@/types/SelectItem";

export type SelectProps = LayoutProps & {
    items: SelectItem[],
    selectedItem: string,
    title: string,
    additionalItems?: SelectItem[],
    setSelectedItem?: (item: string) => void,
    selectedAdditionalItem?: string,
    setSelectedAdditionalItem?: (item: string) => void,
};
