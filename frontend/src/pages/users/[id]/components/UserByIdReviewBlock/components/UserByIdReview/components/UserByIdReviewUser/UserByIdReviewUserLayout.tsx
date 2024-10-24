import {LayoutProps} from "@/types/LayoutProps";

const UserByIdReviewUserLayout = (props: LayoutProps) => {
    const {children} = props;

    return (
        <div className={"flex items-center gap-4 text-base font-semibold"}>
            {children}
        </div>
    )
}

export default UserByIdReviewUserLayout;