def has_add_permission(self, request, obj=None):
     # Not too much elegant but works to hide show_save_and_add_another button
    if '/change/' in str(request):
         return False 
    return True

def get_readonly_fields(self, request, obj=None):
    if obj: # editing an existing object
        # All model fields as read_only
        return self.readonly_fields + tuple([item.name for item in obj._meta.fields])
    return self.readonly_fields


class ReadOnlyModelAdmin(admin.ModelAdmin):
    """ModelAdmin class that prevents modifications through the admin.

    The changelist and the detail view work, but a 403 is returned
    if one actually tries to edit an object.
    """

    actions = None

    def get_readonly_fields(self, request, obj=None):
        return self.fields or [f.name for f in self.model._meta.fields]

    def has_add_permission(self, request):
        return False

    # Allow viewing objects but not actually changing them
    def has_change_permission(self, request, obj=None):
        if request.method not in ('GET', 'HEAD'):
            return False
        return super(ReadOnlyModelAdmin, self).has_change_permission(request, obj)

    def has_delete_permission(self, request, obj=None):
        return False