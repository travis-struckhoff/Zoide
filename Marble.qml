import QtQuick 2.0


Item {
    id: marble

    property int type: 3
    
    property bool exists: false


    Image {
        id: img

        anchors.fill: parent
        source: {
            if (type == 1)
                return "images/redStone.png"
            else if (type == 2)
                return "images/blueStone.png"
            else if (type == 3)
                return "images/yellowStone.png"
            else
                return "images/greenStone.png"
        }
        opacity: 0.1
    }

    states: [
        State {
            name: "ExistsState"
            when: exists == true
            PropertyChanges { target: img; opacity: 1 }
            PropertyChanges { target: img; source: {
            	if (type == 1)
	                 "images/redStone.png"
	            else if (type == 2)
	                 "images/blueStone.png"
	            else if (type == 3)
	                 "images/yellowStone.png"
	            else
	                 "images/greenStone.png"
            	}
            }
        },
        State {
            name: "emptyState"
            when: exists == false
            PropertyChanges { target: img; opacity: 0.1 }
            PropertyChanges { target: img; source: 
                "images/yellowStone.png"}
        }
    ]
}