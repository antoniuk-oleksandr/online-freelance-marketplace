import {LayoutProps} from "@/types/LayoutProps";

const RightSignSideLayout = (props: LayoutProps) => {
    const {children} = props;

    return (
        <div className={"size-full hidden lg:block"}>
            {children}
        </div>
    )
}

export default RightSignSideLayout;