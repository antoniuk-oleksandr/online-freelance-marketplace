import {LayoutProps} from "@/types/LayoutProps";
import {motion} from "framer-motion";
import animation from "@/utils/animation.json";

const NotFoundLayout = (props: LayoutProps) => {
    const {children} = props;

    return (
        <motion.div
            {...animation}
            className={"min-h-app flex flex-col font-semibold text-light-palette-text-primary dark:text-dark-palette-text-primary select-none justify-center items-center"}
        >
            {children}
        </motion.div>
    )
}

export default NotFoundLayout;