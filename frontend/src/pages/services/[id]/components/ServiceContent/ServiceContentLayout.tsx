import {LayoutProps} from "@/types/LayoutProps";

const ServiceContentLayout = (props: LayoutProps) => {
    const {children} = props;

    return (
        <div className={"grid grid-cols-1 lg:grid-cols-service gap-8 mt-8"}>
            {children}
        </div>
    )
}

export default ServiceContentLayout;