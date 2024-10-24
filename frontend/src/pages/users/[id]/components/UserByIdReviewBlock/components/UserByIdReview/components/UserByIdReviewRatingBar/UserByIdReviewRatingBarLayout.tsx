import {LayoutProps} from "@/types/LayoutProps";

const UserByIdReviewRatingBarLayout = (props: LayoutProps) => {
    const {children} = props;

    return (
        <div className={"flex gap-x-2 items-center"}>
            {children}
        </div>
    )
}

export default UserByIdReviewRatingBarLayout;