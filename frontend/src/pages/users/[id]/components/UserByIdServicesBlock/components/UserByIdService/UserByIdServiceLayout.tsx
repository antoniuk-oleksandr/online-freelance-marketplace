import {LayoutProps} from "@/types/LayoutProps";
import Link from "next/link";
import {UserService} from "@/types/UserService";
import PaperElement from "@/common-components/PaperElement";

type UserByIdServiceLayoutProps = LayoutProps & UserService;

const UserByIdServiceLayout = (props: UserByIdServiceLayoutProps) => {
    const {children, id} = props;

    return (
        <Link href={`/services/${id}`}>
            <PaperElement>
                {children}
            </PaperElement>
        </Link>
    )
}

export default UserByIdServiceLayout;