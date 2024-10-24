import ServiceTopLeftSideContentLayout from "./ServiceTopLeftSideContentLayout";
import UserInfoBlock from "@/pages/users/[id]/components/UserInfoBlock/UserInfoBlock";
import DividerElement from "@/common-components/DividerElement";
import ServiceImageCarousel from "@/pages/services/[id]/components/ServiceImageCarousel/ServiceImageCarousel";
import {Service} from "@/types/Service";

type ServiceLeftSideContentProps = Service;

const ServiceTopLeftSideContent = (props: ServiceLeftSideContentProps) => {
    const {title, freelancer} = props;

    return (
        <ServiceTopLeftSideContentLayout>
            <h2 className={"text-2xl font-semibold mb-4"}>{title}</h2>
            <UserInfoBlock
                useLink={true}
                showLanguages={false}
                size={"small"}
                user={freelancer}
            />
            <DividerElement className={"my-8"}/>
            <ServiceImageCarousel {...props}/>
        </ServiceTopLeftSideContentLayout>
    )
}

export default ServiceTopLeftSideContent;