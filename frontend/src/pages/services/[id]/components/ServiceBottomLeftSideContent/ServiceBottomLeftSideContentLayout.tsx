import {LayoutProps} from "@/types/LayoutProps";

const ServiceBottomLeftSideContentLayout = (props: LayoutProps) => {
    const {children} = props;

    return (
        <div className={"flex flex-col gap-y-8"}>
            {children}
        </div>
    )
}

export default ServiceBottomLeftSideContentLayout;