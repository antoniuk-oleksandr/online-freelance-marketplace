import {UserByIdProps} from "@/types/UserByIdProps";
import {Spoiler} from "@mantine/core";
import UserComponentLayout from "@/pages/users/[id]/components/UserComponentLayout";

const UserAboutBlock = ({user}: UserByIdProps) => {
    const {about} = user;

    if(!about) return null;
    return (
        <UserComponentLayout>
            <h2 className="text-xl font-bold">About</h2>
            <Spoiler maxHeight={72} showLabel="Show more" hideLabel="Hide">
                <p
                    dangerouslySetInnerHTML={{__html: about}}
                    className="text-base text-light-palette-text-secondary dark:text-dark-palette-text-secondary"/>
            </Spoiler>
        </UserComponentLayout>
    );
}

export default UserAboutBlock;
