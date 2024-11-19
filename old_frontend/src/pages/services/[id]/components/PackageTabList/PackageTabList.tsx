import PackageTabListLayout from "./PackageTabListLayout";
import {Package} from "@/types/Package";
import PackageTab from "@/pages/services/[id]/components/PackageTab/PackageTab";
import {Dispatch, SetStateAction} from "react";

type PackageTabListProps = {
    packages: Package[],
    tabs: string[],
    selectedTab: number,
    setSelectedTab: Dispatch<SetStateAction<number>>
}

const PackageTabList = (props: PackageTabListProps) => {
    const {packages, tabs} = props;

    return (
        <PackageTabListLayout>
            {packages.map((_, index) => (
                <PackageTab
                    key={index}
                    {...props}
                    index={index}
                    length={packages.length}
                    text={tabs[index]}
                />
            ))}
        </PackageTabListLayout>
    )
}

export default PackageTabList;