<script lang="ts">
    import {createForm} from "felte";
    import type {SignProps} from "@/types/SignProps.ts";
    import {validator} from "@felte/validator-zod";
    import {formStore} from "@/common-components/Sign/stores/form-store.js";
    import type {FormStore} from "@/types/FormStore.ts";
    import {useLocation} from "svelte-routing";

    const {children, submitAction, schema, defaultValues}: SignProps = $props();

    const {form, data, errors, isSubmitting, reset, setFields} = createForm({
        initialValues: defaultValues,
        extend: validator({schema}),
        onSubmit: async (values) => submitAction(values)
    })

    useLocation().subscribe((_) => {
        formStore.set({data: null, errors: null, wasSubmitted: false, keepSignedIn: false});
        reset();
    })

    $effect(() => formStore.update((prev: FormStore) => ({...prev, data: $data})))
    $effect(() => formStore.update((prev: FormStore) => ({...prev, errors: $errors})))
    $effect(() => {
        if ($isSubmitting) {
            formStore.update((prev: FormStore) => ({...prev, wasSubmitted: true}))
        }
    })

    let prevKeepSignedIn = $state(false);
    formStore.subscribe((value) => {
        if(value.keepSignedIn !== prevKeepSignedIn) {
            prevKeepSignedIn = value.keepSignedIn;
            setFields('keepSignedIn', value.keepSignedIn);
        }
    })
</script>

<form
        class="flex flex-col gap-y-5"
        use:form>
    {#if children}
        {@render children()}
    {/if}
</form>