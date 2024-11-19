import {LayoutProps} from "@/types/LayoutProps";
import {useTheme} from "next-themes";
import {Dispatch, SetStateAction} from "react";

type PackageTabLayoutProps = LayoutProps & {
    length: number,
    index: number,
    setSelectedTab: Dispatch<SetStateAction<number>>,
    selectedTab: number
}

const PackageTabLayout = (props: PackageTabLayoutProps) => {
    const {children, length, index, selectedTab, setSelectedTab} = props;
    const width = length === 1 ? "w-full" : length === 2 ? "w-1/2" : "w-1/3";
    const {resolvedTheme} = useTheme();
    const unselectedColor = resolvedTheme === 'dark' ? 'rgba(255, 255, 255, 0.12)' : 'rgba(0, 0, 0, 0.12)';
    const selectedColor = 'rgb(6, 182, 212)';
    const selected = index === selectedTab;

    return (
        <div
            onClick={() => setSelectedTab(props.index)}
            style={{borderColor: selected ? selectedColor : unselectedColor}}
            className={`${width} cursor-pointer hover:bg-transparent !p-4 !text-base !font-semibold text-center border-b`}
        >
            {children}
        </div>
    )
}

export default PackageTabLayout;