import {LayoutProps} from "@/types/LayoutProps";
import {motion} from "framer-motion";
import animation from "@/utils/animation.json"

const UserPageLayout = (props: LayoutProps) => {
    const {children} = props;

    return (
        <motion.div
            {...animation}
            className={"flex flex-col gap-y-8"}>
            {children}
        </motion.div>
    )
}

export default UserPageLayout;