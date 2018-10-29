# Generated by Django 2.0.6 on 2018-06-30 12:01

from django.db import migrations
from django.db.migrations import RunPython


def fnc(apps, schema_editor):
    Module = apps.get_model('mydiscord', 'Module')
    Module.objects.get_or_create(name='publicroles')


class Migration(migrations.Migration):
    dependencies = [
        ('publicroles', '0001_initial'),
        ('mydiscord', '0001_initial'),
    ]

    operations = [
        RunPython(fnc, RunPython.noop)
    ]