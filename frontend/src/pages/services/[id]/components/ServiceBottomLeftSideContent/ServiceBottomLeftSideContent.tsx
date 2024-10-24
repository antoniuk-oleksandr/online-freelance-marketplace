import {Service} from "@/types/Service";
import ServiceBottomLeftSideContentLayout from "./ServiceBottomLeftSideContentLayout";
import ServiceAboutBlock from "@/pages/services/[id]/components/ServiceAboutBlock/ServiceAboutBlock";
import UserByIdReviewBlock from "@/pages/users/[id]/components/UserByIdReviewBlock/UserByIdReviewBlock";

type ServiceBottomLeftSideContentProps = Service;

const ServiceBottomLeftSideContent = (props: ServiceBottomLeftSideContentProps) => {
    const {reviews} = props;

    return (
        <ServiceBottomLeftSideContentLayout>
            <ServiceAboutBlock {...props}/>
            <UserByIdReviewBlock showServices={false} reviews={reviews}/>
        </ServiceBottomLeftSideContentLayout>
    )
}

export default ServiceBottomLeftSideContent;