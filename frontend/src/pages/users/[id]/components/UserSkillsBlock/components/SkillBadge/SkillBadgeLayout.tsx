import {LayoutProps} from "@/types/LayoutProps";

const SkillBadgeLayout = (props: LayoutProps) => {
    const {children} = props;

    return (
        <div className={"rounded-xl border border-light-palette-divider dark:border-dark-palette-divider px-3 py-1 bg-light-palette-background-block dark:bg-dark-palette-background-block"}>
            {children}
        </div>
    )
}

export default SkillBadgeLayout;