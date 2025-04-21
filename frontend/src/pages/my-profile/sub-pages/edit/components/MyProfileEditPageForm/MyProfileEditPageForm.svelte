<script lang="ts">
  import type { FilterParams } from '@/types/FilterParams'
  import type { LayoutProps } from '@/types/LayoutProps'
  import type { MyProfileEditPageFormData } from '@/types/MyProfileEditPageFormData'
  import { createForm } from 'felte'
  import { setContext } from 'svelte'
  import { handleMyProfileEditPageFormSubmit } from '../../handlers'
  import { validator } from '@felte/validator-zod'
  import { myProfileEditPageFormScheme } from '../../helpers'

  type MyProfileEditPageFormProps = LayoutProps & {
    data: MyProfileEditPageFormData
    filterParams: FilterParams
  }

  const { data: initialValues, children }: MyProfileEditPageFormProps = $props()

  const { form, data, isSubmitting, errors } = createForm<MyProfileEditPageFormData>({
    initialValues,
    onSubmit: (data) => handleMyProfileEditPageFormSubmit(data, initialValues),
    extend: validator({ schema: myProfileEditPageFormScheme }),
  })

  setContext('formDataStore', data)
  setContext('formErrorsStore', errors)
  setContext('formIsSubmittingStore', isSubmitting)
</script>

<form class="flex flex-col md:grid md:grid-cols-edit-profile-form gap-6 flex-grow" use:form>
  {@render children()}
</form>
