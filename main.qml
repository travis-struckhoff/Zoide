/****************************************************************************
**
** Copyright (C) 2013 Digia Plc and/or its subsidiary(-ies).
** Contact: http://www.qt-project.org/legal
**
** This file is part of the Qt Quick Controls module of the Qt Toolkit.
**
** $QT_BEGIN_LICENSE:BSD$
** You may use this file under the terms of the BSD license as follows:
**
** "Redistribution and use in source and binary forms, with or without
** modification, are permitted provided that the following conditions are
** met:
**   * Redistributions of source code must retain the above copyright
**     notice, this list of conditions and the following disclaimer.
**   * Redistributions in binary form must reproduce the above copyright
**     notice, this list of conditions and the following disclaimer in
**     the documentation and/or other materials provided with the
**     distribution.
**   * Neither the name of Digia Plc and its Subsidiary(-ies) nor the names
**     of its contributors may be used to endorse or promote products derived
**     from this software without specific prior written permission.
**
**
** THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
** "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
** LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
** A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
** OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
** SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
** LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
** DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
** THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
** (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
** OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE."
**
** $QT_END_LICENSE$
**
****************************************************************************/

import QtQuick 2.0
import "../Zoide" as Z

Rectangle {
    id: screen

    width: 400; height: 400

    SystemPalette { id: activePalette }

    Rectangle {
        id: topArrows
        width: parent.width; height: 30
        color: activePalette.window
        anchors.top: screen.top

        Z.Button {
            id: topLeft
            property int arrow: 0
            anchors { left: parent.left; verticalCenter: parent.verticalCenter }
            text: "Rotate Left"
            onClicked: game.handleArrows(this)
        }
        Z.Button {
            id: topLeft2
            property int arrow: 1
            anchors { left: topLeft.right; verticalCenter: parent.verticalCenter }
            text: "Rotate Right"
            onClicked: game.handleArrows(this)
        }
        Z.Button {
            id: topRight
            property int arrow: 3
            anchors { right: parent.right; verticalCenter: parent.verticalCenter }
            text: "Rotate Right"
            onClicked: game.handleArrows(this)
        }
        Z.Button {
            id: topRight2
            property int arrow: 2
            anchors { right: topRight.left; verticalCenter: parent.verticalCenter }
            text: "Rotate Left"
            onClicked: game.handleArrows(this)
        }
    }

    Rectangle {
        id: bottomArrows
        width: parent.width; height: 30
        color: activePalette.window
        anchors.bottom: toolBar.top

        Z.Button {
            id: bottomLeft
            property int arrow: 4
            anchors { left: parent.left; verticalCenter: parent.verticalCenter }
            text: "Rotate Left"
            onClicked: game.handleArrows(this)
        }
        Z.Button {
            id: bottomLeft2
            property int arrow: 5
            anchors { left: bottomLeft.right; verticalCenter: parent.verticalCenter }
            text: "Rotate Right"
            onClicked: game.handleArrows(this)
        }
        Z.Button {
            id: bottomRight
            property int arrow: 7
            anchors { right: parent.right; verticalCenter: parent.verticalCenter }
            text: "Rotate Right"
            onClicked: game.handleArrows(this)
        }
        Z.Button {
            id: bottomRight2
            property int arrow: 6
            anchors { right: bottomRight.left; verticalCenter: parent.verticalCenter }
            text: "Rotate Left"
            onClicked: game.handleArrows(this)
        }
    }

    Item {
        width: parent.width
        anchors { top: topArrows.bottom; bottom: bottomArrows.top }

        Image {
            id: background
            anchors.fill: parent
            source: "images/space.jpg"
            fillMode: Image.PreserveAspectCrop
        }

        Item {
            id: gameCanvas

            property int score: 0
            property int blockSize: 40

            width: parent.width - (parent.width % blockSize)
            height: parent.height - (parent.height % blockSize)
            anchors.centerIn: parent

            MouseArea {
                anchors.fill: parent
                onClicked: game.handleClick(mouse.x, mouse.y)
            }
        }
    }

    Z.Dialog {
        id: dialog
        anchors.centerIn: parent
        z: 100
    }

    Rectangle {
        id: toolBar
        width: parent.width; height: 30
        color: activePalette.window
        anchors.bottom: screen.bottom

        Z.Button {
            anchors { left: parent.left; verticalCenter: parent.verticalCenter }
            text: "New Game"
            onClicked: game.startNewGame(gameCanvas, dialog)
        }

        Text {
            objectName: "score"
            id: score
            anchors { right: parent.right; verticalCenter: parent.verticalCenter }
            text: "Score: Who knows?"
        }
    }
}
