import {LayoutProps} from "@/types/LayoutProps";

const UserByIdReviewBottomContentLayout = (props: LayoutProps) => {
    const {children} = props;

    return (
        <div className={"flex flex-col gap-4"}>
            {children}
        </div>
    )
}

export default UserByIdReviewBottomContentLayout;