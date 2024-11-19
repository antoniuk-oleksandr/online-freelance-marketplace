import {LayoutProps} from "@/types/LayoutProps";
import PageLayout from "@/common-components/PageLayout";

const FooterLayout = (props: LayoutProps) => {
    const {children} = props;

    return (
        <footer
            className={"border bg-light-palette-background-block dark:bg-dark-palette-background-block border-light-palette-divider dark:border-dark-palette-divider py-4"}>
            <PageLayout>
                <div className={"flex md:flex-row flex-col gap-y-4 justify-between items-center text-light-palette-text-secondary dark:text-dark-palette-text-secondary"}>
                    {children}
                </div>
            </PageLayout>
        </footer>
    )
}

export default FooterLayout;