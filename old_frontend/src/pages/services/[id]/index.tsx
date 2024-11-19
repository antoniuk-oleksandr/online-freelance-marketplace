import {useServiceById} from "@/pages/services/[id]/hooks/use-service-by-id";
import NotFound from "@/common-components/NotFound/NotFound";
import AppLoader from "@/common-components/AppLoader";
import ServiceBreadcrumbs from "@/pages/services/[id]/components/ServiceBreadcrumbs/ServiceBreadcrumbs";
import ServicePageLayout from "@/pages/services/[id]/ServicePageLayout";
import ServicePackagesBlock from "@/pages/services/[id]/components/ServicePackagesBlock/ServicePackagesBlock";
import ServiceTopLeftSideContent from "@/pages/services/[id]/components/ServiceTopLeftSideContent/ServiceTopLeftSideContent";
import ServiceBottomLeftSideContent
    from "@/pages/services/[id]/components/ServiceBottomLeftSideContent/ServiceBottomLeftSideContent";
import ServiceContent from "@/pages/services/[id]/components/ServiceContent/ServiceContent";

const ServicePage = () => {
    const {service, status} = useServiceById();

    if (status !== 200 && status !== null) return <NotFound/>
    if (!service) return <AppLoader/>

    return (
        <ServicePageLayout>
            <ServiceBreadcrumbs {...service}/>
            <ServiceContent {...service}/>
        </ServicePageLayout>
    );
}

export default ServicePage;