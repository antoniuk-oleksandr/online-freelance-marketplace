import {LayoutProps} from "@/types/LayoutProps";
import PageLayout from "@/common-components/PageLayout";

const HeaderLayout = (props: LayoutProps) => {
    const {children} = props;

    return (
        <header className={"py-4 bg-light-palette-background-block dark:bg-dark-palette-background-block border-b border-light-palette-divider dark:border-dark-palette-divider shadow-sm w-full"}>
            <PageLayout>
                <div className={"flex items-center justify-between"}>
                    {children}
                </div>
            </PageLayout>
        </header>
    )
}

export default HeaderLayout;