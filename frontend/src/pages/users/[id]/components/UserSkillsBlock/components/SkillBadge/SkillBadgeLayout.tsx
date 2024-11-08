import {LayoutProps} from "@/types/LayoutProps";
import Link from "next/link";

type SkillBadgeLayoutProps = LayoutProps & {
    id: number
};

const SkillBadgeLayout = (props: SkillBadgeLayoutProps) => {
    const {children, id} = props;

    return (
        <Link
            href={"/search?skills=" + id}
            className={"rounded-xl border border-light-palette-divider dark:border-dark-palette-divider px-3 py-1 bg-light-palette-background-block dark:bg-dark-palette-background-block"}
        >
            {children}
        </Link>
    )
}

export default SkillBadgeLayout;