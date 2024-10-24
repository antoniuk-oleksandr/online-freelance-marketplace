import PackageTabLayout from "./PackageTabLayout";
import {Dispatch, SetStateAction} from "react";

type PackageTabProps = {
    text: string,
    length: number,
    setSelectedTab: Dispatch<SetStateAction<number>>,
    selectedTab: number,
    index: number,
}

const PackageTab = (props: PackageTabProps) => {
    const {text} = props;

    return (
        <PackageTabLayout {...props}>
            <span>{text}</span>
        </PackageTabLayout>
    )
}

export default PackageTab;