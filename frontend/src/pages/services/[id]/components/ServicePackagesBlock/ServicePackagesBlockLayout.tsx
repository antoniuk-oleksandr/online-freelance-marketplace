import {LayoutProps} from "@/types/LayoutProps";
import PaperElement from "@/common-components/PaperElement";

const ServicePackagesBlockLayout = (props: LayoutProps) => {
    const {children} = props;

    return (
        <PaperElement
            styles={"row-start-2 !p-0 col-start-1 md:row-start-1 md:col-start-2 border border-light-palette-divider dark:border-dark-palette-divider h-fit top-packages md:sticky"}
        >
            {children}
        </PaperElement>
    )
}

export default ServicePackagesBlockLayout;