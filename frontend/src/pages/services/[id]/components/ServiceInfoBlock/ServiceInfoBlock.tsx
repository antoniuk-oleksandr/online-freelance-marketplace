import ServiceInfoBlockLayout from "./ServiceInfoBlockLayout";
import UserInfoBlock from "@/pages/users/[id]/components/UserInfoBlock/UserInfoBlock";
import DividerElement from "@/common-components/DividerElement";
import {Service} from "@/types/Service";
import ServiceImageCarousel from "@/pages/services/[id]/components/ServiceImageCarousel/ServiceImageCarousel";
import ServiceAboutBlock from "@/pages/services/[id]/components/ServiceAboutBlock/ServiceAboutBlock";

type ServiceTypeProps = Service;

const ServiceInfoBlock = (props: ServiceTypeProps) => {
    const {title, freelancer} = props;

    return (
        <ServiceInfoBlockLayout>
            123
        </ServiceInfoBlockLayout>
    )
}

export default ServiceInfoBlock;