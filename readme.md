# ETD Cli

This is the cli for ETD admin. Use this to download installation template.

## Usage

### Create a .env file locally

You need to create an env file along with your cli executable. An example .env file is below.
```
etd_node_id=DEVICE_UNIQUE_ID
```

### Run the cli tool

Available flags are

**Template**: Template tag for installation template

**Environment**: Any of beta, production, local

**Password**: Remote admin password

### Sample command

```bash
./cli --template=etdnet --environment=beta
```