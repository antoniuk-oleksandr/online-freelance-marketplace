import {LayoutProps} from "@/types/LayoutProps";

const UserByIdReviewServiceImageLayout = (props: LayoutProps) => {
    const {children} = props;

    return (
        <div className={"aspect-video h-9 w-auto"}>
            {children}
        </div>
    )
}

export default UserByIdReviewServiceImageLayout;