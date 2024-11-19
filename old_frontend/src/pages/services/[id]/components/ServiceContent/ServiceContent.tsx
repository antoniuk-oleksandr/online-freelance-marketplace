import ServiceContentLayout from "./ServiceContentLayout";
import ServiceTopLeftSideContent
    from "@/pages/services/[id]/components/ServiceTopLeftSideContent/ServiceTopLeftSideContent";
import ServiceBottomLeftSideContent
    from "@/pages/services/[id]/components/ServiceBottomLeftSideContent/ServiceBottomLeftSideContent";
import ServicePackagesBlock from "@/pages/services/[id]/components/ServicePackagesBlock/ServicePackagesBlock";
import {Service} from "@/types/Service";

type ServiceTypeProps = Service;

const ServiceContent = (props: ServiceTypeProps) => {
    const {packages} = props;

    return (
        <ServiceContentLayout>
            <ServiceTopLeftSideContent {...props}/>
            <ServiceBottomLeftSideContent {...props}/>
            <ServicePackagesBlock packages={packages}/>
        </ServiceContentLayout>
    )
}

export default ServiceContent;