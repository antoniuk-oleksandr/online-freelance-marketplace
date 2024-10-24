import {LayoutProps} from "@/types/LayoutProps";

const ServiceBreadcrumbsLayout = (props: LayoutProps) => {
    const {children} = props;

    return (
        <div className={"mb-8"}>
            {children}
        </div>
    )
}

export default ServiceBreadcrumbsLayout;