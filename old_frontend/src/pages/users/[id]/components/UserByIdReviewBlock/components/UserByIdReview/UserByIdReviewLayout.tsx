import {LayoutProps} from "@/types/LayoutProps";
import PaperElement from "@/common-components/PaperElement";

const UserByIdReviewLayout = (props: LayoutProps) => {
    const {children} = props;

    return (
        <PaperElement styles={"flex flex-col gap-4"}>
            {children}
        </PaperElement>
    )
}

export default UserByIdReviewLayout;