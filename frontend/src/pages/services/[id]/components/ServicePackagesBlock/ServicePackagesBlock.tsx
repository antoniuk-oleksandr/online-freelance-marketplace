import ServicePackagesBlockLayout from "./ServicePackagesBlockLayout";
import {Package} from "@/types/Package";
import PackageTabList from "@/pages/services/[id]/components/PackageTabList/PackageTabList";
import {useState} from "react";
import SelectedPackage from "@/pages/services/[id]/components/SelectedPackage/SelectedPackage";

type ServicePackagesBlockProps = {
    packages: Package[]
}

const ServicePackagesBlock = (props: ServicePackagesBlockProps) => {
    const {packages} = props;
    const tabs = ['Basic', 'Standard', 'Premium'];
    const [selectedTab, setSelectedTab] = useState(0);

    return (
        <ServicePackagesBlockLayout>
            <PackageTabList
                setSelectedTab={setSelectedTab}
                selectedTab={selectedTab}
                packages={packages}
                tabs={tabs}
            />
            <SelectedPackage {...packages[selectedTab]}/>
        </ServicePackagesBlockLayout>
    )
}

export default ServicePackagesBlock;