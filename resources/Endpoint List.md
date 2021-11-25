# Exhaustive Endpoint List
This is an exhaustive list of endpoints for reference 


POST / Create
------------
Perform an HTTP POST operation with the appropriate JSON 
```
/api
/api/login
/api/tenants
/api/sites
/api/buildings
/api/rooms
/api/acs
/api/walls
/api/panels
/api/aisles
/api/tiles
/api/cabinets
/api/groups
/api/corridors
/api/racksensors
/api/devicesensors
/api/racks
/api/devices
/api/room-templates
/api/obj-templates
```



DELETE / Delete
------------
Perform an HTTP DELETE operation without JSON body
```
/api/tenants/{id}
/api/sites/{id}
/api/sites
/api/buildings/{id}
/api/rooms/{id}
/api/rooms/{id}/acs/{id}
/api/rooms/{id}/panels/{id}
/api/rooms/{id}/walls/{id}
/api/rooms/{id}/aisles/{id}
/api/rooms/{id}/tiles/{id}
/api/rooms/{id}/cabinets/{id}
/api/rooms/{id}/groups/{id}
/api/rooms/{id}/corridors/{id}
/api/racks/{id}/racksensors/{id}
/api/devices/{id}/devicesensors/{id}
/api/racks/{id}
/api/devices/{id}
/api/room-templates/{template_name}
/api/obj-templates/{template_name}
```


PUT / Update
-------------
Perform an HTTP PUT operation with desired JSON body
```
/api/tenants/{id}
/api/sites/{id}
/api/buildings/{id}
/api/rooms/{id}
/api/rooms/{id}/acs/{id}
/api/rooms/{id}/panels/{id}
/api/rooms/{id}/walls/{id}
/api/rooms/{id}/aisles/{id}
/api/rooms/{id}/tiles/{id}
/api/rooms/{id}/cabinets/{id}
/api/rooms/{id}/groups/{id}
/api/rooms/{id}/corridors/{id}
/api/racks/{id}/racksensors/{id}
/api/devices/{id}/devicesensors/{id}
/api/racks/{id}
/api/devices/{id}
/api/room-templates/{template_name}
/api/obj-templates/{template_name}
```

PATCH / Update
-------------
Perform an HTTP PUT operation with desired JSON body
```
/api/tenants/{id}
/api/sites/{id}
/api/buildings/{id}
/api/rooms/{id}
/api/rooms/{id}/acs/{id}
/api/rooms/{id}/panels/{id}
/api/rooms/{id}/walls/{id}
/api/rooms/{id}/aisles/{id}
/api/rooms/{id}/tiles/{id}
/api/rooms/{id}/cabinets/{id}
/api/rooms/{id}/groups/{id}
/api/rooms/{id}/corridors/{id}
/api/racks/{id}/racksensors/{id}
/api/devices/{id}/devicesensors/{id}
/api/racks/{id}
/api/devices/{id}
/api/room-templates/{template_name}
/api/obj-templates/{template_name}
```

GET / Get
-------------
Perform an HTTP PUT operation without JSON

### Quick Token Check
This URL is for development purposes only
```
/api/token/valid
```

### Get All Objects
```
/api/tenants
/api/sites
/api/buildings
/api/rooms
/api/racks
/api/devices
/api/room-templates
/api/obj-templates
```

### Get by ID (non hierarchal)
ID is a long string    
Template_name is the 'slug'
```
/api/tenants/{id}
/api/sites/{id}
/api/buildings/{id}
/api/rooms/{id}
/api/racks/{id}
/api/devices/{id}
/api/room-templates/{template_name}
/api/obj-templates/{template_name}
```

### Search Objects
Search using a query in the URL
Objects that match the query will be returned
Example: /devices?name=myValue?color=silver
will return silver devices with name 'myValue'
```
/api/acs?
/api/walls?
/api/panels?
/api/tenants?
/api/sites?
/api/buildings?
/api/rooms?
/api/racks?
/api/devices?
/api/aisles?
/api/tiles?
/api/cabinets?
/api/groups?
/api/corridors?
/api/racksensors?
/api/devicesensors?
/api/room-templates?
/api/obj-templates?
```

### Objects Only Obtainable via hierarchal GET
These objects cannot be obtained directly and must have their hierarchy specified in the URL. If the URL does not have {id}, then all objects under the given hierarchy will be returned
```
/api/rooms/{id}/acs
/api/rooms/{id}/panels
/api/rooms/{id}/walls
/api/rooms/{id}/acs/{id}
/api/rooms/{id}/panels/{id}
/api/rooms/{id}/walls/{id}
/api/rooms/{id}/aisles
/api/rooms/{id}/tiles
/api/rooms/{id}/cabinets
/api/rooms/{id}/groups
/api/rooms/{id}/corridors
/api/rooms/{id}/aisles/{id}
/api/rooms/{id}/tiles/{id}
/api/rooms/{id}/cabinets/{id}
/api/rooms/{id}/groups/{id}
/api/rooms/{id}/corridors/{id}
/api/racks/{id}/racksensors/{id}
/api/devices/{id}/devicesensors/{id}
```

### Get all Objects 2 levels lower
```
/api/tenants/{tenant_name}/buildings
/api/sites/{id}/rooms
/api/buildings/{id}/racks
/api/rooms/{id}/devices
```

### Get an Object's entire hierarchy
The object and everything related to it will be returned    
in a nested JSON fashion
```
/api/tenants/{tenant_name}/all
/api/sites/{id}/all
/api/buildings/{id}/all
/api/rooms/{id}/all
/api/racks/{id}/all
/api/devices/{id}/all
```

### Get object's ranged hierarchy 
Limits the depth of the hierarchy to retrieve. This is observed by the    
URL given. 
```
/api/tenants/{tenant_name}/all/sites/buildings/rooms/racks/devices
/api/tenants/{tenant_name}/all/sites/buildings/rooms/racks
/api/tenants/{tenant_name}/all/sites/buildings/rooms
/api/tenants/{tenant_name}/all/sites/buildings

/api/sites/{id}/all/buildings/rooms/racks/devices
/api/sites/{id}/all/buildings/rooms/racks
/api/sites/{id}/all/buildings/rooms

/api/buildings/{id}/all/rooms/racks/devices
/api/buildings/{id}/all/rooms/racks

/api/rooms/{id}/all/racks/devices
```

### Get objects through the hierarchy
Returns an object if name given or all the objects immediately under the given URL
```
/api/tenants/{tenant_name}/sites/{site_name}/buildings/{building_name}/rooms/{room_name}/racks/{rack_name}/devices/{device_name}
/api/tenants/{tenant_name}/sites/{site_name}/buildings/{building_name}/rooms/{room_name}/racks/{rack_name}/devices
/api/tenants/{tenant_name}/sites/{site_name}/buildings/{building_name}/rooms/{room_name}/racks/{rack_name}
/api/tenants/{tenant_name}/sites/{site_name}/buildings/{building_name}/rooms/{room_name}/racks
/api/tenants/{tenant_name}/sites/{site_name}/buildings/{building_name}/rooms/{room_name}
/api/tenants/{tenant_name}/sites/{site_name}/buildings/{building_name}/rooms
/api/tenants/{tenant_name}/sites/{site_name}/buildings/{building_name}
/api/tenants/{tenant_name}/sites/{site_name}/buildings
/api/tenants/{tenant_name}/sites/{site_name}
/api/tenants/{tenant_name}/sites



/api/sites/{id}/buildings/{building_name}/rooms/{room_name}/racks/{rack_name}/devices/{device_name}
/api/sites/{id}/buildings/{building_name}/rooms/{room_name}/racks/{rack_name}/devices
/api/sites/{id}/buildings/{building_name}/rooms/{room_name}/racks/{rack_name}
/api/sites/{id}/buildings/{building_name}/rooms/{room_name}/racks
/api/sites/{id}/buildings/{building_name}/rooms/{room_name}
/api/sites/{id}/buildings/{building_name}/rooms
/api/sites/{id}/buildings/{building_name}
/api/sites/{id}/buildings


/api/buildings/{id}/rooms/{room_name}/racks/{rack_name}/devices/{device_name}
/api/buildings/{id}/rooms/{room_name}/racks/{rack_name}/devices
/api/buildings/{id}/rooms/{room_name}/racks/{rack_name}
/api/buildings/{id}/rooms/{room_name}/racks
/api/buildings/{id}/rooms/{room_name}
/api/buildings/{id}/rooms


/api/rooms/{id}/racks/{rack_name}/devices/{device_name}
/api/rooms/{id}/racks/{rack_name}/devices
/api/rooms/{id}/racks/{rack_name}
/api/rooms/{id}/racks


/api/racks/{id}/devices/{device_name}
/api/racks/{id}/devices
```

### Get object's hierarchy (non standard)
This returns an object's hierarchy in a non standard fashion    
and will be removed in the future
```
/api/tenants/{tenant_name}/all/nonstd
/api/sites/{id}/all/nonstd
/api/buildings/{id}/all/nonstd
/api/rooms/{id}/all/nonstd
/api/racks/{id}/all/nonstd
```