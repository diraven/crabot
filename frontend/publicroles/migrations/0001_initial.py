# Generated by Django 2.0.6 on 2018-06-17 10:52

from django.db import migrations, models
import django.db.models.deletion


class Migration(migrations.Migration):

    initial = True

    dependencies = [
        ('mydiscord', '0002_alias'),
    ]

    operations = [
        migrations.CreateModel(
            name='PublicRole',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('uid', models.IntegerField(unique=True)),
                ('guild', models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, related_name='public_roles', to='mydiscord.Guild')),
            ],
        ),
    ]