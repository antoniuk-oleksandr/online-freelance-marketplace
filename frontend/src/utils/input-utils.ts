import {FieldErrors} from "react-hook-form";

export const getErrorMessage = (
    errors: FieldErrors,
    id: string,
) => {
    return errors[id] ? typeof errors[id].message === 'string' ? errors[id].message : '' : '';
}