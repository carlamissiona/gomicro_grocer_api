# Grocer Service

This is the Grocer service

Generated with

```
micro new grocer
```

## Usage

Generate the proto code

```
make proto
```

Run the service

```
micro run .
```

## List of Endpoints
**Greetings**
**GetGroceryList**
**Call**
```
micro grocer Geetings  --name="client 1 name" --greetings="Hi " --address="dsds"
micro grocer GetGroceryList  --name="client 1 name" --greetings="Hi " --address="dsds"
micro grocer Call  --name="client 1 name" --greetings="Hi " --address="dsds"
```
