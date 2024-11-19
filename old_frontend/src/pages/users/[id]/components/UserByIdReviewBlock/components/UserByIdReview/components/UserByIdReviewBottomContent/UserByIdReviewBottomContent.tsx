import UserByIdReviewBottomContentLayout from "./UserByIdReviewBottomContentLayout";
import DividerElement from "@/common-components/DividerElement";
import {getPackageDuration} from "@/utils/utils";
import UserByIdReviewContent
    from "@/pages/users/[id]/components/UserByIdReviewBlock/components/UserByIdReview/components/UserByIdReviewContent";
import {Review} from "@/types/Review";
import UserByIdReviewColumn
    from "@/pages/users/[id]/components/UserByIdReviewBlock/components/UserByIdReview/components/UserByIdReviewColumn/UserByIdReviewColumn";
import UserByIdReviewService
    from "@/pages/users/[id]/components/UserByIdReviewBlock/components/UserByIdReview/components/UserByIdReviewService/UserByIdReviewService";

type UserByIdReviewBottomContentProps = {
    review: Review,
    showServices?: boolean
}

const UserByIdReviewBottomContent = (props: UserByIdReviewBottomContentProps) => {
    const {review, showServices} = props;
    const {createdAt, endedAt, service} = review;

    console.log(review);

    const {price} = service;

    return (
        <UserByIdReviewBottomContentLayout>
            <UserByIdReviewContent {...review}/>
            <div className={"flex gap-y-1 md:gap-x-8 md:h-12 md:items-center flex-col md:flex-row items-start"}>
                <UserByIdReviewColumn label={"Price"} data={`$${price}`}/>
                <DividerElement className={"hidden md:block"} orientation={"vertical"}/>
                <UserByIdReviewColumn label={"Duration"} data={getPackageDuration(createdAt, endedAt)}/>
                {showServices === false
                    ? null
                    : <>
                        <DividerElement className={"hidden md:block"} orientation={"vertical"}/>
                        <p className={"text-light-palette-text-secondary dark:text-dark-palette-text-secondary mt-2 md:hidden"}>Ordered</p>
                        <UserByIdReviewService {...service}/>
                    </>
                }

            </div>
        </UserByIdReviewBottomContentLayout>
    )
}

export default UserByIdReviewBottomContent;