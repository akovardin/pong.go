//
//  Controller.swift
//  Pongmobile
//
//  Created by Artem Kovardin on 10.10.2021.
//

import Foundation
import UIKit
import SwiftUI

class GameController: MobileEbitenViewController {
    func onErrorOnGameUpdate(err: NSError) {
        print(err)
    }
        
}

struct GameView: UIViewControllerRepresentable {

    // 2.
    func makeUIViewController(context: Context) -> MobileEbitenViewController {
        return MobileEbitenViewController()
    }
    
    // 3.
    func updateUIViewController(_ uiViewController: MobileEbitenViewController, context: Context) {
        
    }
}
