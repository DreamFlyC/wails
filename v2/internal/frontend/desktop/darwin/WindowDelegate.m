//
//  WindowDelegate.m
//  test
//
//  Created by Lea Anthony on 10/10/21.
//

#import <Foundation/Foundation.h>
#import <Cocoa/Cocoa.h>
#import "WindowDelegate.h"

@implementation WindowDelegate

- (BOOL)windowShouldClose:(NSWindow *)sender {
    [sender orderOut:nil];
    if( self.hideOnClose == false ) {
        NSLog(@"send message: WC");
    }
    return !self.hideOnClose;
}

@end